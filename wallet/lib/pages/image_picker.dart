import 'dart:io';

import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:image_cropper/image_cropper.dart';
import 'package:image_picker/image_picker.dart';
import 'package:pylons_wallet/components/loading.dart';
import 'package:pylons_wallet/model/pick_image_model.dart';
import 'package:pylons_wallet/services/repository/repository.dart';
import 'package:pylons_wallet/utils/constants.dart';

Future<File?> pickImageFromGallery(double maxHeight, double maxWidth,
    int imageQuality, BuildContext context) async {
  final pickImageEither = await GetIt.I.get<Repository>().pickImageFromGallery(
      PickImageModel(
          maxHeight: maxHeight,
          maxWidth: maxWidth,
          imageQuality: imageQuality,
          imageSource: ImageSource.gallery));

  if (pickImageEither.isLeft()) {
    // ignore: use_build_context_synchronously
    pickImageEither.swap().toOption().toNullable()!.message.show();
    return null;
  }
  final pickedImagePath = pickImageEither.getOrElse(() => '');

  if (pickedImagePath.isEmpty) {
    return null;
  }

  return cropImage(pickedImagePath);
}

Future<File?> cropImage(String path) async {
  return ImageCropper.cropImage(
      sourcePath: path,
      aspectRatioPresets: [
        CropAspectRatioPreset.square,
        CropAspectRatioPreset.ratio3x2,
        CropAspectRatioPreset.original,
        CropAspectRatioPreset.ratio4x3,
        CropAspectRatioPreset.ratio16x9
      ],
      androidUiSettings: const AndroidUiSettings(
          toolbarTitle: kStripeMerchantDisplayName,
          toolbarColor: kBlue,
          toolbarWidgetColor: Colors.white,
          initAspectRatio: CropAspectRatioPreset.original,
          lockAspectRatio: false),
      iosUiSettings: const IOSUiSettings(
        minimumAspectRatio: 1.0,
      ));
}
