#import "AppDelegate.h"
#import "LoginViewController.h"

@interface AppDelegate ()

@end

@implementation AppDelegate

- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
        self.window = [[UIWindow alloc] initWithFrame:UIScreen.mainScreen.bounds];

        UIStoryboard *storyboard = [UIStoryboard storyboardWithName:@"Main" bundle:nil];
        LoginViewController *loginVc = [storyboard instantiateViewControllerWithIdentifier:@"LoginViewControllerID"];
        UINavigationController *nav = [[UINavigationController alloc] initWithRootViewController:loginVc];
        self.window.rootViewController = nav;
        [self.window makeKeyWindow];
        return YES;
}

- (void)applicationWillTerminate:(UIApplication *)application {
    if (self.isScreenShareing) {
        [MeetingControlCompoments endShareScreen];
        [[MeetingRTCManager shareRtc] stopScreenSharing];
    }
}

#pragma mark - UISceneSession lifecycle

- (UIInterfaceOrientationMask)application:(UIApplication *)application supportedInterfaceOrientationsForWindow:(UIWindow *)window {
    if (self.allowAutoRotate) {
        return UIInterfaceOrientationMaskLandscape | UIInterfaceOrientationMaskPortrait;
    } else {
        //支持竖屏
        //Support vertical screen
        return UIInterfaceOrientationMaskPortrait;
    }
}


@end
