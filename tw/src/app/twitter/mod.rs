
use app::config::{Config, TwitterConfig, UserConfig};
use failure::{Error, bail, format_err};
use std::io;
use std::io::Write;
use std::borrow::Cow;
use tokio_core::reactor::Core;
use webbrowser;
use egg_mode::{Token};

impl UserConfig {
    pub fn access(&self) -> KeyPair {
        KeyPair{
            key: self.access_token.clone(),
            secret: self.access_token_secret.clone(),
        }
    }

    pub fn to_egg(self) -> Token {
        Token::Access{
            consumer: egg_mode::KeyPair{
                key: Cow::from(self.consumer_key),
                secret: Cow::from(self.consumer_secret),
            },
            access: egg_mode::KeyPair{
                key: Cow::from(self.access_token),
                secret: Cow::from(self.access_token_secret),
            },
        }
    }
}

#[derive(Serialize, Deserialize, Clone)]
pub struct KeyPair {
    pub key: String,
    pub secret: String,
}

impl KeyPair {
    pub fn create_from_setting(twitter_info: String) -> KeyPair {
        let twitter_value: Vec<&str> = twitter_info.as_str().split("\n").collect();
        let consumer_key = twitter_value[0].to_string();
        let consumer_secret = twitter_value[1].to_string();
        KeyPair {
            key: consumer_key,
            secret: consumer_secret,
        }
    }

    pub fn to_egg(self) -> egg_mode::KeyPair {
        egg_mode::KeyPair {
            key: Cow::from(self.key),
            secret: Cow::from(self.secret),
        }
    }
}

pub struct Twitter {
    consumer: KeyPair,
    access: KeyPair,
    pub user_id: u64,
    pub screen_name: String,
}

impl Twitter {
    pub fn create_from_config(config: &Config, username: String, consumer: &KeyPair) -> Result<Twitter, Error>  {
        let user_cfg = match config.users.get(&username) {
            Some(cfg) => cfg,
            None => bail!("user not found {}", username),
        };

        let consumer = Twitter::select_keypair(&config.twitter, consumer);
        let access = user_cfg.access();

        Ok(Twitter {
            consumer: consumer,
            access: access,
            user_id: 0,
            screen_name: String::from(""),
        })
    }

    pub fn new(cfg: &TwitterConfig, consumer: &KeyPair) -> Result<Twitter, Error> {
        let cfg_key = Twitter::select_keypair(&cfg, &consumer);

        let mut core = Core::new()?;
        let handle = core.handle();
        let con_token = egg_mode::KeyPair::new(consumer.key.clone(), consumer.secret.clone());
        let request_token = core.run(egg_mode::request_token(&con_token, "oob", &handle))?;

        let auth_url = egg_mode::authorize_url(&request_token);
        webbrowser::open(auth_url.as_str())?;

        print!("input pin: ");
        io::stdout().flush().unwrap();
        let mut pin = String::new();
        io::stdin().read_line(&mut pin)?;

        let (token, user_id, screen_name) = core.run(egg_mode::access_token(
            con_token,
            &request_token,
            pin,
            &handle,
        ))?;
        let (access_token, access_token_secret) = match token {
            egg_mode::Token::Access {
                consumer: _,
                access:
                    egg_mode::KeyPair {
                        key: ak,
                        secret: ats,
                    },
            } => (ak, ats),
            _ => bail!("unexpected token"),
        };
        Ok(Twitter {
            consumer: cfg_key,
            access: KeyPair{
                key: access_token.to_string(),
                secret: access_token_secret.to_string(),
            },
            user_id: user_id,
            screen_name: screen_name,
        })
    }

    pub fn keys(&self) -> (KeyPair, KeyPair) {
        return (self.consumer.clone(), self.access.clone())
    }

    fn select_keypair(cfg: &TwitterConfig, consumer: &KeyPair) -> KeyPair {
        let cfg_key = KeyPair {
            key: cfg.consumer_key.clone(),
            secret: cfg.consumer_secret.clone(),
        };
        if cfg.consumer_key == "" || cfg.consumer_secret == "" {
            consumer.clone()
        } else {
            cfg_key.clone()
        }
    }
}
