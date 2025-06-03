import 'package:get/get_navigation/src/routes/get_route.dart';
import 'package:hecto_front/routes/route.dart';

class HAppRoute {
  static final List<GetPage> pages = [
    GetPage(name: HRoute.login, page: () => AuthenticationPage()),
    GetPage(name: HRoute.game, page: () => GameSelctor()),
    GetPage(name: HRoute.playGame, page: () => PlayGame()),
    GetPage(name: HRoute.home, page: () => HomePage()),
  ];
}
