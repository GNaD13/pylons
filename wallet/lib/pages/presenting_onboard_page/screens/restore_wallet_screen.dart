import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter_svg/svg.dart';
import 'package:get_it/get_it.dart';
import 'package:pylons_wallet/pages/presenting_onboard_page/components/import_from_google_form.dart';
import 'package:pylons_wallet/stores/wallet_store.dart';
import 'package:pylons_wallet/utils/svg_util.dart';

class RestoreWalletScreen extends StatelessWidget {
  const RestoreWalletScreen({Key? key}) : super(key: key);

  WalletsStore get walletsStore => GetIt.I.get();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Positioned(
            left: 0,
            right: 0,
            top: 0,
            bottom: 0,
            child: SvgPicture.asset(
              SVGUtil.CREATE_WALLET_BACKGROUND,
              width: MediaQuery.of(context).size.width,
              height: MediaQuery.of(context).size.height,
              fit: BoxFit.cover,
            ),
          ),
          Positioned(
              left: 0,
              right: 0,
              top: 0,
              bottom: 0,
              child: SingleChildScrollView(
                child: Padding(
                  padding: EdgeInsets.symmetric(horizontal: 25.w),
                  child: ImportFromGoogleForm(
                    walletStore: walletsStore,
                  ),
                ),
              )),
        ],
      ),
    );
  }
}
