import 'package:flutter/material.dart';
import 'package:pecs_mobile/Screens/ActionPage/action_page.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/model/category.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import '../../../constants.dart';
import '../home_page.dart';

final storage = FlutterSecureStorage();

class CategoryListTileWidget extends StatelessWidget {
  final Category category;

  Image imageFromBase64String(String base64String) {
    return Image.memory(base64Decode(base64String),
        height: 120, width: 80, fit: BoxFit.cover);
  }

  Future<String> deleteCategory(String id) async {
    var jwt = await storage.read(key: "jwt");

    var res = await http.delete(
      "$SERVER_IP/category/" + id,
      headers: {"Authorization": jwt},
    );

    if (res.statusCode == 200) {
      return "ciao";
    }
    print(res.body);
    print(res.headers);
    print(res.statusCode);
    return null;
  }

  const CategoryListTileWidget({
    Key key,
    @required this.category,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final style = TextStyle(fontSize: 18);

    return GestureDetector(
      onTap: () {
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) {
              return ActionPage(categoryId: category.id, categoryName: category.name);
            },
          ),
        );
      },
      onLongPress: () {
        _buildPopupDialog(context, category);
      },
      child: Column(
        children: <Widget>[
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              Expanded(
                child: Column(
                  children: <Widget>[
                    SizedBox(height: 2),
                    imageFromBase64String(category.img),
                    ListTile(
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) {
                              return ActionPage(categoryId: category.id, categoryName: category.name);
                            },
                          ),
                        );
                      },
                      //leading:
                      title: Text(
                        category.name,
                        style: style,
                      ),
                      trailing: Icon(Icons.arrow_forward),
                    ),
                  ],
                ),
              )
            ],
          ),
        ],
      ),
    );
  }

  Future<void> _buildPopupDialog(BuildContext context, Category category) {
    return showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text("Vuoi eliminare la categoria?"),
            content: SingleChildScrollView(
              child: ListBody(
                children: <Widget>[
                  RoundedButton(
                    text: "SI",
                    color: kPrimaryLightColor,
                    textColor: Colors.black,
                    press: () async {
                      var res = await deleteCategory(category.id);
                      if (res != null) {
                        Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => HomePage()));                      } else {
                        displayDialog(context, "Error!", "ERRORE");
                        //TODO controll other errors
                      }
                    },
                  ),
                  RoundedButton(
                    text: "NO",
                    color: kPrimaryLightColor,
                    textColor: Colors.black,
                    press: () {
                      Navigator.of(context).pop();
                    },
                  ),
                ],
              ),
            ),
          );
        });
  }

  void displayDialog(context, title, text) => showDialog(
    context: context,
    builder: (context) =>
        AlertDialog(
            title: Text(title),
            content: Text(text)
        ),
  );
}
