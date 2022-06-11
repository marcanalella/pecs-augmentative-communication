import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/components/rounded_input_field.dart';
import 'package:pecs_mobile/constants.dart';
import 'package:pecs_mobile/model/category.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:image_picker/image_picker.dart';

import 'category_listtile.dart';
import 'home_page_background.dart';
import 'package:http/http.dart' as http;

final storage = FlutterSecureStorage();

class HomePageBody extends StatefulWidget {
  @override
  _HomePageBodyState createState() => _HomePageBodyState();
}

Future<String> insertCategory(String name, String img) async {
  Map data = {"name": name, "img": img};

  var body = json.encode(data);
  var jwt = await storage.read(key: "jwt");

  var res = await http.post("$SERVER_IP/category",
      headers: {"Authorization": jwt}, body: body);

  if (res.statusCode == 201) {
    return "ciao"; //TODO
  }
  print(res.body);
  print(res.headers);
  print(res.statusCode);
  return null;
}

class _HomePageBodyState extends State<HomePageBody> {
  List<Category> categories;
  final TextEditingController _titleController = TextEditingController();

  Future<List<Category>> getCategories() async {
    var jwt = await storage.read(key: "jwt");
    //if (jwt == null) return "";
    print(jwt);
    var res = await http
        .get("$SERVER_IP/categories", headers: {"Authorization": jwt});

    if (res.statusCode == 200) {
      categories = (json.decode(res.body) as List)
          .map((i) => Category.fromJson(i))
          .toList();
      return categories;
    }

    print(res.body);
    print(res.headers);
    print(res.statusCode);

    return null;
    //TODO Handle errors
  }

  File imageFile;

  _openGallery(BuildContext context) async {
    var img = await ImagePicker.pickImage(source: ImageSource.gallery);
    this.setState(() {
      imageFile = img;
    });
  }

  _openCamera(BuildContext context) async {
    var img = await ImagePicker.pickImage(source: ImageSource.camera);
    this.setState(() {
      imageFile = img;
    });
  }

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    // This size provide us total height and width of our screen
    return HomePageBackground(
      child: FutureBuilder(
          future: getCategories(),
          builder: (context, snapshot) => (categories != null &&
                  categories.length != 0)
              ? Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    SizedBox(height: size.height * 0.05),
                    Text(
                      "CIAO, GIORGIO!",
                      style: TextStyle(fontWeight: FontWeight.bold),
                    ),
                    SizedBox(height: size.height * 0.06),
                    Expanded(
                      child: GridView.count(
                        crossAxisCount: 2,
                        shrinkWrap: true,
                        children: categories.map((item) {
                          return CategoryListTileWidget(
                            category: item,
                          );
                        }).toList(),
                      ),
                    ),
                    SizedBox(height: 24),
                    Container(
                      padding:
                          EdgeInsets.symmetric(horizontal: 32, vertical: 12),
                      color: Colors.transparent,
                      child: RoundedButton(
                        text: "+ CATEGORIA",
                        color: kPrimaryLightColor,
                        textColor: Colors.black,
                        press: () {
                          _showChoiceDialog(context);
                        },
                      ),
                    )
                  ],
                )
              : snapshot.hasError
                  ? Text("Errore durante il caricamento")
                  : (categories != null && categories.length == 0)
                      ? Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                              Text("Lista vuota"),
                              RoundedButton(
                                text: "+ CATEGORIA",
                                color: kPrimaryLightColor,
                                textColor: Colors.black,
                                press: () {
                                  _showChoiceDialog(context);
                                },
                              )
                            ])
                      : CircularProgressIndicator()),
    );
  }

  void displayDialog(context, title, text) => showDialog(
        context: context,
        builder: (context) =>
            AlertDialog(title: Text(title), content: Text(text)),
      );

  Future<void> _showChoiceDialog(BuildContext context) {
    return showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text("Inserisci categoria"),
            content: SingleChildScrollView(
              child: ListBody(
                children: <Widget>[
                  Text("Inserisci titolo e immagine"),
                  RoundedInputField(
                      hintText: "Titolo",
                      emailController: _titleController,
                      icon: Icons.arrow_forward),
                  SizedBox(height: 5),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      ElevatedButton(
                        onPressed: () {
                          _openGallery(context);
                        },
                        child:
                            Icon(Icons.perm_media_sharp, color: Colors.black),
                        style: ElevatedButton.styleFrom(
                          shape: CircleBorder(),
                          padding: EdgeInsets.all(20),
                          primary: kPrimaryLightColor, // <-- Button color
                        ),
                      ),
                      ElevatedButton(
                        onPressed: () {
                          _openCamera(context);
                        },
                        child: Icon(Icons.photo_camera, color: Colors.black),
                        style: ElevatedButton.styleFrom(
                          shape: CircleBorder(),
                          padding: EdgeInsets.all(20),
                          primary: kPrimaryLightColor, // <-- Button color
                        ),
                      )
                    ],
                  ),
                  imageShow(),
                  RoundedButton(
                    text: "Aggiungi",
                    press: () async {
                      var name = _titleController.text;
                      var img = base64Encode(imageFile.readAsBytesSync());
                      var jwt = await insertCategory(name, img);
                      if (jwt != null) {
                        setState(() {});
                        Navigator.of(context).pop();
                        displayDialog(context, "OK!", "Categoria inserita correttamente");
                      } else {
                        displayDialog(context, "Error!", "Errore nel caricamento della categoria");
                      }
                    },
                  )
                ],
              ),
            ),
          );
        });
  }

  Widget imageShow() {
    if (imageFile != null) {
      return Column(
        children: [
          SizedBox(height: 10),
          Image.file(
            imageFile,
            fit: BoxFit.cover,
          )
        ],
      );
    }
    return SizedBox(height: 0.1);
  }
}
