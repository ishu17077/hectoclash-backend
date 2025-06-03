import 'package:flutter/src/widgets/navigator.dart';
import 'package:get/get_navigation/src/routes/route_middleware.dart';

class RoutesMiddleware extends GetMiddleware{
  @override
  RouteSettings? redirect(String? route) {
    
    return super.redirect(route);
  }
}
