import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/components/rounded_input_field.dart';
import 'package:pecs_mobile/constants.dart';
import 'package:pecs_mobile/model/action.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:image_picker/image_picker.dart';

import 'action_listtile.dart';
import 'action_page_background.dart';
import 'package:http/http.dart' as http;

final storage = FlutterSecureStorage();

class ActionPageBody extends StatefulWidget {
  final String categoryId;
  final String categoryName;

  ActionPageBody({
    Key key,
    @required this.categoryId,
    @required this.categoryName})
      : super(key: key);

  @override
  _ActionPageBodyState createState() => _ActionPageBodyState();
}

class _ActionPageBodyState extends State<ActionPageBody> {
  List<Ations> actions;
  final TextEditingController _titleController = TextEditingController();

  Future<String> insertAction(String name, String img) async {
    Map data = {"name": name, "img": img, "categoryId": widget.categoryId};

    var body = json.encode(data);
    var jwt = await storage.read(key: "jwt");

    var res = await http.post("$SERVER_IP/action",
        headers: {"Authorization": jwt}, body: body);

    if (res.statusCode == 201) {
      return "ciao"; //TODO
    }
    print(res.body);
    print(res.headers);
    print(res.statusCode);
    return null;
  }

  Future<List<Ations>> getActions() async {
    var jwt = await storage.read(key: "jwt");
    //if (jwt == null) return "";
    print(jwt);
    var res = await http.get("$SERVER_IP/actions/" + widget.categoryId,
        headers: {"Authorization": jwt});

    if (res.statusCode == 200) {
      actions = (json.decode(res.body) as List)
          .map((i) => Ations.fromJson(i))
          .toList();
      return actions;
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
    return ActionPageBackground(
        child: FutureBuilder(
      future: getActions(),
      builder: (context, snapshot) => (actions != null && actions.length != 0)
          ? Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                SizedBox(height: size.height * 0.05),
                Text(
                  widget.categoryName,
                  style: TextStyle(fontWeight: FontWeight.bold),
                ),
                SizedBox(height: size.height * 0.06),
                Expanded(
                  child: GridView.count(
                    crossAxisCount: 2,
                    shrinkWrap: true,
                    children: actions.map((item) {
                      return ActionListTileWidget(
                        action: item,
                      );
                    }).toList(),
                  ),
                ),
                SizedBox(height: 24),
                Container(
                  padding: EdgeInsets.symmetric(horizontal: 32, vertical: 12),
                  color: Colors.transparent,
                  child: RoundedButton(
                    text: "+ AZIONE",
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
                  : (actions != null && actions.length == 0)
                      ? Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                              Text("Lista vuota"),
                              RoundedButton(
                                text: "+ AZIONE",
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
            title: Text("Inserisci Azione"),
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
                      var jwt = await insertAction(name, img);
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
