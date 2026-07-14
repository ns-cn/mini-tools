import 'package:flutter/material.dart';
import 'package:get/get.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(),
    );
  }
}

class Controller extends GetxController {
  var count = 0.obs;
  increment() => count++;
}

class MyHomePage extends StatelessWidget {
  const MyHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    final Controller controller = Get.put(Controller());
    return Column(children: [
      const Spacer(),
      const Text("Windows"),
      const Spacer(),
      Row(
        children: [
          IconButton(
            icon: const Icon(
              Icons.window,
              size: 32,
            ),
            onPressed: controller.increment,
          ),
          const Spacer(),
          Obx(() => Text("${controller.count}")),
        ],
      )
    ]);
  }
}
