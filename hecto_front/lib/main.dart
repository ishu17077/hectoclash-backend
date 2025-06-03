import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:hecto_front/routes/app_route.dart';
import 'package:hecto_front/routes/route.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'HectoClash',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.white60),
      ),
      initialRoute: HRoute.home,
      getPages: HAppRoute.pages,
      unknownRoute: GetPage(
        name: '/404',
        page: () => const Scaffold(body: Center(child: Text("404 Not Found"))),
      ),
      // home: const HectoClash(),
    );
  }
}
