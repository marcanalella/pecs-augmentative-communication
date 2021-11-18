import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:pecs_mobile/Screens/HomePage/home_page.dart';
import 'package:pecs_mobile/Screens/Login/components/login_background.dart';
import 'package:pecs_mobile/components/already_have_an_account_acheck.dart';
import 'package:pecs_mobile/components/rounded_button.dart';
import 'package:pecs_mobile/components/rounded_input_field.dart';
import 'package:pecs_mobile/components/rounded_password_field.dart';
import 'package:pecs_mobile/model/login_response.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_svg/svg.dart';
import 'package:flutter_web_browser/flutter_web_browser.dart';
import 'package:http/http.dart' as http;
import 'package:url_launcher/url_launcher.dart';

import '../../../constants.dart';

final storage = FlutterSecureStorage();

class LoginBody extends StatelessWidget {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  Future<String> attemptLogIn(String username, String password) async {
    Map data = {"email": username, "password": password};

    var body = json.encode(data);

    var res = await http.post("$SERVER_IP/login", body: body);

    if (res.statusCode == 200) {
      Map<String, dynamic> responseJson = json.decode(res.body);
      var loginResponse = LoginResponse.fromJson(responseJson);
      return loginResponse.accessToken;
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Background(
      child: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              "BENVENUTO!",
              style: TextStyle(fontWeight: FontWeight.bold),
            ),
            SizedBox(height: size.height * 0.03),
            SvgPicture.asset(
              "assets/icons/chat.svg",
              height: size.height * 0.35,
            ),
            SizedBox(height: size.height * 0.03),
            RoundedInputField(
              hintText: "Email",
              emailController: _usernameController,
            ),
            RoundedPasswordField(
              passwordController: _passwordController,
            ),
            RoundedButton(
              text: "LOGIN",
              press: () async {
                var username = _usernameController.text;
                var password = _passwordController.text;
                var jwt = await attemptLogIn(username, password);
                if (jwt != null) {
                  storage.write(key: "jwt", value: jwt);
                  Navigator.push(context,
                      MaterialPageRoute(builder: (context) => HomePage()));
                } else {
                  displayDialog(context, "Error!",
                      "No account was found matching that username and password.");
                  //TODO controll other errors
                }
              },
            ),
            SizedBox(height: size.height * 0.03),
            AlreadyHaveAnAccountCheck(
              press: () async {
                await FlutterWebBrowser.openWebPage(url: url);
              },
            ),
          ],
        ),
      ),
    );
  }

  void displayDialog(context, title, text) => showDialog(
        context: context,
        builder: (context) =>
            AlertDialog(title: Text(title), content: Text(text)),
      );
}
