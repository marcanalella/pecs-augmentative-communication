import 'package:flutter/material.dart';
import 'package:pecs_mobile/components/text_field_container.dart';
import 'package:pecs_mobile/constants.dart';

class RoundedPasswordField extends StatelessWidget {
  final ValueChanged<String> onChanged;
  final TextEditingController passwordController;
  const RoundedPasswordField({
    Key key,
    this.onChanged,
    this.passwordController
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFieldContainer(
      child: TextField(
        obscureText: true,
        cursorColor: kPrimaryColor,
        controller: passwordController,
        decoration: InputDecoration(
          hintText: "Password",
          icon: Icon(
            Icons.lock,
            color: kPrimaryColor,
          ),
          suffixIcon: Icon(
            Icons.visibility,
            color: kPrimaryColor,
          ),
          border: InputBorder.none,
        ),
      ),
    );
  }
}
