import 'package:flutter/material.dart';
import 'package:pecs_mobile/constants.dart';

class AlreadyHaveAnAccountCheck extends StatelessWidget {
  final Function press;
  const AlreadyHaveAnAccountCheck({
    Key key,
    this.press,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Text(
          "Non hai ancora un account? ",
          style: TextStyle(color: kPrimaryColor),
        ),
        GestureDetector(
          onTap: press,
          child: Text(
            "Registrati",
            style: TextStyle(
              color: kPrimaryColor,
              fontWeight: FontWeight.bold,
            ),
          ),
        )
      ],
    );
  }
}
