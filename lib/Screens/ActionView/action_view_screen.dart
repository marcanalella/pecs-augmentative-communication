import 'package:flutter/material.dart';
import 'package:pecs_mobile/model/action.dart';

import 'components/action_view_body.dart';

class ActionViewScreen extends StatelessWidget {
  final Ations action;

  const ActionViewScreen({Key key, @required this.action}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ActionViewBody(action: action),
    );
  }
}
