import 'package:flutter/material.dart';
import 'package:pecs_mobile/components/text_field_container.dart';
import 'package:pecs_mobile/constants.dart';

class RoundedInputField extends StatelessWidget {
  final String hintText;
  final IconData icon;
  final ValueChanged<String> onChanged;
  final TextEditingController emailController;

  const RoundedInputField({
    Key key,
    this.hintText,
    this.icon,
    this.onChanged,
    this.emailController
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFieldContainer(
      child: TextField(
        controller: emailController,
        cursorColor: kPrimaryColor,
        decoration: InputDecoration(
          icon: Icon(
            icon,
            color: kPrimaryColor,
          ),
          hintText: hintText,
          border: InputBorder.none,
        ),
      ),
    );
  }
}
