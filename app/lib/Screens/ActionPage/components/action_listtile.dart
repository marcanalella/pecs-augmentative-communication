import 'package:flutter/material.dart';
import 'package:pecs_mobile/Screens/ActionView/action_view_screen.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/model/action.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import '../../../constants.dart';
import '../action_page.dart';

final storage = FlutterSecureStorage();

class ActionListTileWidget extends StatelessWidget {
  final Ations action;

  Image imageFromBase64String(String base64String) {
    return Image.memory(base64Decode(base64String),
        height: 120, width: 80, fit: BoxFit.cover);
  }

  Future<String> deleteCategory(String id) async {
    var jwt = await storage.read(key: "jwt");

    var res = await http.delete(
      "$SERVER_IP/action/" + id,
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

  const ActionListTileWidget({
    Key key,
    @required this.action,
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
              return ActionViewScreen(action: action);
            },
          ),
        );
      },
      onLongPress: () {
        _buildPopupDialog(context, action);
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
                    imageFromBase64String(action.img),
                    ListTile(
                      onTap: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) {
                              return ActionViewScreen(action: action);
                            },
                          ),
                        );
                      },
                      //leading:
                      title: Text(
                        action.name,
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

  Future<void> _buildPopupDialog(BuildContext context, Ations action) {
    return showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text("Vuoi eliminare l\'azione?"),
            content: SingleChildScrollView(
              child: ListBody(
                children: <Widget>[
                  RoundedButton(
                    text: "SI",
                    color: kPrimaryLightColor,
                    textColor: Colors.black,
                    press: () async {
                      var res = await deleteCategory(action.id);
                      if (res != null) {
                        Navigator.pop(
                            context,
                            MaterialPageRoute(
                                builder: (context) =>
                                    ActionPage(categoryId: action.categoryId)));
                      } else {
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
            AlertDialog(title: Text(title), content: Text(text)),
      );
}
