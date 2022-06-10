import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:pecs_mobile/constants.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'Screens/HomePage/home_page.dart';
import 'Screens/Login/login_screen.dart';

final storage = FlutterSecureStorage();

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  Future<String> get jwtOrEmpty async {
    var jwt = await storage.read(key: "jwt");
    if (jwt == null) return "";
    return jwt;
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Auth',
      theme: ThemeData(
        primaryColor: kPrimaryColor,
        scaffoldBackgroundColor: Colors.white,
      ),
      home: FutureBuilder(
          future: jwtOrEmpty,
          builder: (context, snapshot) {
            if (!snapshot.hasData) return CircularProgressIndicator();
            if (snapshot.data != "") {
              var str = snapshot.data;
              var jwt = str.split(".");

              if (jwt.length != 3) {
                return LoginScreen();
              } else {
                var payload = json.decode(
                    ascii.decode(base64.decode(base64.normalize(jwt[1]))));
                if (DateTime.fromMillisecondsSinceEpoch(payload["exp"] * 1000)
                    .isAfter(DateTime.now())) {
                  return HomePage();
                } else {
                  return LoginScreen();
                }
              }
            } else {
              return LoginScreen();
            }
          }),
    );
  }
}
