class Category {
  String id;
  String name;
  String img;

  Category({this.id, this.name, this.img});

  Category.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        img = json['img'];

  Map<String, dynamic> toJson() => {
    'id': id,
    'name': name,
    'img': img,
  };
}