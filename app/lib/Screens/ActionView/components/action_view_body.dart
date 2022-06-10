import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:pecs_mobile/Screens/HomePage/home_page.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/model/action.dart';
import 'package:flutter_tts/flutter_tts.dart';

import 'action_view_background.dart';
import 'action_view_divider.dart';
import 'action_view_share_icon.dart';

FlutterTts flutterTts = FlutterTts();

class ActionViewBody extends StatelessWidget {
  final Ations action;

  const ActionViewBody({Key key, @required this.action}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return ActionViewBackground(
      child: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              action.name,
              style: TextStyle(fontWeight: FontWeight.bold),
            ),
            SizedBox(height: size.height * 0.03),
            imageFromBase64String(action.img, size.height),
            SizedBox(height: size.height * 0.03),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                ShareIcon(
                  iconSrc: "assets/icons/speaker.svg",
                  press: () async {
                    await flutterTts.setLanguage("it-IT");
                    await flutterTts.setPitch(1);
                    await flutterTts.speak(action.name);
                  },
                ),
              ],
            ),
            ActionViewDivider(),
            RoundedButton(
              text: "INDIETRO",
              press: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) {
                      return HomePage();
                    },
                  ),
                );
              },
            ),
          ],
        ),
      ),
    );
  }

  Image imageFromBase64String(String base64String, double height) {
    return Image.memory(
      base64Decode(base64String),
      height: height * 0.35,
    );
  }
}
