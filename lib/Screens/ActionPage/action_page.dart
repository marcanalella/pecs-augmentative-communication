import 'package:flutter/material.dart';

import 'components/action_page_body.dart';

class ActionPage extends StatelessWidget {
  final String categoryId;
  final String categoryName;

  const ActionPage({Key key, @required this.categoryId, @required this.categoryName}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ActionPageBody(categoryId: categoryId, categoryName: categoryName,),
    );
  }
}
