import 'package:products/services/products_service.dart';
import 'package:shelf_router/shelf_router.dart';
import 'package:shelf/shelf.dart';
import 'package:shelf/shelf_io.dart' as io;

void main(List<String> arguments) async {
  final app = Router();

  app.mount('/products/', ProductService().router);

  app.get('/hello', (Request req) {
    return Response.ok('Hello World!!!');
  });

  await io.serve(app, 'localhost', 1323);
}
