import 'dart:convert';
import 'package:shelf_router/shelf_router.dart';
import 'package:shelf/shelf.dart';
import 'package:products/models/products.dart';
import 'package:json_annotation/src/allowed_keys_helpers.dart';

class ProductService {
  final data = ProductList();

  Router get router {
    final router = Router();

    router.get('/', (Request req)  => okResponse(json.encode(data.products)));

    router.get('/<id|[0-9]+>', (Request req, String id) {
      try {
        final parsedId = int.parse(id);
        final product = data.getProduct(parsedId);
        return okResponse(json.encode(product));
      } catch (e) {
        return notFound(json.encode({'message': (e as FormatException).message}));
      }
    });

    router.put('/<id|[0-9]+>', (Request req, String id) async {
      try {
        final parsedId = int.parse(id);
        final payload = await req.readAsString();
        final newProduct = Product.fromJson(json.decode(payload));

        data.updateProduct(newProduct, parsedId);
        return Response.ok(payload,
            headers: {'Content-Type': 'application/json'});
      } catch (e) {
        if (e is FormatException) {
          return notFound(json.encode({'message': e.message}));
        } 
        else if (e is BadKeyException) {
          return badRequest(json.encode({'message': e.message}));
        }
        return Response.internalServerError();
      }
    });

    router.post('/', (Request req) async {
      try {
        final payload = await req.readAsString();
        final product = Product.fromJson(json.decode(payload));
        data.addProduct(product);
        return Response(201, body: 'product created');
      } catch (error) {
        if (error is BadKeyException) {
          return badRequest(json.encode({'message': error.message}));
        }

        return Response.internalServerError();
      }
    });

    router.delete('/<id|[0-9]+>', (Request req, String id) {
      try {
        final parsedId = int.parse(id);
        data.deleteProduct(parsedId);
        return Response.ok('Product deleted');
      } catch(e) {
        if (e is FormatException) {
          return notFound(json.encode({'message': e.message}));
        }
        return Response.internalServerError();
      }
    });

    return router;
  }

  Response okResponse(String body) => 
    Response.ok(body, headers: {'Content-Type': 'application/json'});

  Response notFound(String body) => 
    Response.notFound(body, headers: {'Content-Type': 'application/json'});

  Response badRequest(String body) => 
    Response(400, body: body, headers: {'Content-Type': 'application/json'});
}
