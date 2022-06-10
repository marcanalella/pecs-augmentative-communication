class Ations {
  String id;
  String name;
  String img;
  String categoryId;


  Ations({this.id, this.name, this.img, this.categoryId});

  Ations.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        img = json['img'],
        categoryId = json['categoryId'];

  Map<String, dynamic> toJson() => {
    'id': id,
    'name': name,
    'img': img,
    'categoryId' : categoryId,
  };
}