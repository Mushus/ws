#[macro_use]
extern crate serde_derive;
extern crate app_dirs;
extern crate clap;
extern crate egg_mode;
extern crate failure;
extern crate tokio_core;
extern crate webbrowser;
extern crate colored;

mod app;

use app::{App, KeyPair};
use app::config::ConfigLoader;
use app_dirs::AppInfo;

const APP_INFO: AppInfo = AppInfo {
    name: "tw",
    author: "Mushus",
};

const TWITTER_APPLICATION: &'static str = include_str!("../twitter_application");

fn main() {
    let app = clap::App::new("tw")
        .version("0.1.0")
        .about("Command Line Twitter Client");

    app.get_matches();

    let mut cl = match ConfigLoader::new(APP_INFO) {
        Ok(cl) => cl,
        Err(err) => return println!("failed to load twitter: {}", err),
    };
    let mut cfg = cl.load();

    let consumer = KeyPair::create_from_setting(TWITTER_APPLICATION.to_string());

    {
        let mut app = App::new(&mut cfg, consumer);
        app.start();
    }

    match cl.save(cfg) {
        Ok(_) => {}
        Err(err) => return println!("failed to save config: {}", err),
    }
}
