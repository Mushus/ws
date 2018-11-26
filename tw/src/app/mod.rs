pub mod config;
mod twitter;

use self::config::{Config, UserConfig};
pub use self::twitter::KeyPair;
use self::twitter::Twitter;
use failure::{Error, bail, format_err};

pub struct App<'a> {
    config: &'a mut Config,
    consumer: KeyPair,
}

impl<'a> App<'a> {
    pub fn new(cfg: &'a mut Config, consumer: KeyPair) -> App<'a> {
        App {
            config: cfg,
            consumer: consumer,
        }
    }

    pub fn start(&mut self) -> Result<(), Error> {
        let tw = match self.config.get_user() {
            Some(uname) => Twitter::create_from_config(self.config, uname, &self.consumer),
            None => Twitter::new(&self.config.twitter, &self.consumer),
        }?;

        if self.timeline(&tw).is_err(){
            println!("error")
        };

        let (consumer, access) = tw.keys();
        self.config.users.insert(
            String::new(),
            UserConfig {
                consumer_key: consumer.key,
                consumer_secret: consumer.secret,
                access_token: access.key,
                access_token_secret: access.secret,
            },
        );
        Ok(())
    }

    fn timeline(&self, tw: &Twitter) -> Result<(), Error> {
        Ok(())
    }
}