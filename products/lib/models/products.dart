import 'package:json_annotation/json_annotation.dart';

part 'products.g.dart';

@JsonSerializable()
class Product {
  @JsonKey(required: true)
  int id;

  @JsonKey(required: true)
  final String name;

  @JsonKey(required: true)
  final String description;

  @JsonKey(required: true)
  final double price;

  @JsonKey(required: true)
  final String sku;

  @JsonKey(name: 'created-at')
  late final String createdAt;

  @JsonKey(name: 'updated-at')
  late String updatedAt;

  @JsonKey(includeIfNull: false, name: 'deleted-at')
  final String? deletedAt;

  Product(
    this.id,
    this.name,
    this.description,
    this.price,
    this.sku,
    this.deletedAt
  ) {
    createdAt = DateTime.now().toString();
    updatedAt = DateTime.now().toString();
  }

  factory Product.fromJson(Map<String, dynamic> json) => _$ProductFromJson(json);
  Map<String, dynamic> toJson() => _$ProductToJson(this);
}

class ProductList {
  final _product_list = <Product>[
    Product(1, 'Latte', 'Frothy milky coffe', 2.45, 'abc123', null),
    Product(2, 'Espresso', 'Short and string coffe without milk', 1.99, 'mgf123', null)
  ];

  List<Product> get products => _product_list;

  void addProduct(Product product) {
    product.id = _product_list.last.id+1;
    _product_list.add(product);
  }

  Product getProduct(int id) {
    final product = _product_list.firstWhere(
      (prod) => prod.id == id, orElse: () => throw FormatException('Product not found')
    );
    
    return product;
  }

  void updateProduct(Product data, int id) {
    final index = _product_list.indexWhere((prod) => prod.id == id);
    if (index == -1) throw FormatException('Product not found');
    _product_list[index] = data;
  }

  void deleteProduct(int id) {
    final index = _product_list.indexWhere((prod) => prod.id == id);
    if (index == -1) throw FormatException('Product not found');
    _product_list.removeAt(index);
  }
}
