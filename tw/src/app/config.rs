use app_dirs::{get_app_root, AppDataType, AppInfo};
use failure::Error;
use std::collections::HashMap;
use std::fs;
use std::fs::{OpenOptions, File};
use std::io::{Read, Write, Seek, SeekFrom};

#[derive(Serialize, Deserialize)]
pub struct Config {
    pub twitter: TwitterConfig,
    pub users: HashMap<String, UserConfig>,
}

impl Config {
    pub fn get_user(&self) -> Option<String> {
        match self.users.keys().next() {
            Some(user) => Some(user.clone()),
            None => None,
        }
    }
}

#[derive(Serialize, Deserialize, Clone)]
pub struct TwitterConfig {
    pub consumer_key: String,
    pub consumer_secret: String,
}

#[derive(Serialize, Deserialize)]
pub struct UserConfig {
    pub consumer_key: String,
    pub consumer_secret: String,
    pub access_token: String,
    pub access_token_secret: String,
}

pub struct ConfigLoader {
    file: File,
    default_twitter: TwitterConfig,
}

impl ConfigLoader {
    pub fn new(app_info: AppInfo) -> Result<ConfigLoader, Error> {
        let app_root = get_app_root(AppDataType::UserConfig, &app_info)?;
        let config_path = app_root.join("config.json");

        //println!("{:?}", app_root);
        fs::create_dir_all(app_root)?;

        let file: File = OpenOptions::new()
            .create(true)
            .read(true)
            .write(true)
            .open(config_path)?;

        Ok(ConfigLoader {
            file: file,
            default_twitter: TwitterConfig {
                consumer_key: String::new(),
                consumer_secret: String::new(),
            },
        })
    }

    fn read(&mut self) -> Result<Config, Error> {
        let mut contents = String::new();
        self.file.read_to_string(&mut contents)?;
        Ok(self.serialize(contents)?)
    }

    fn serialize(&self, contents: String) -> Result<Config, Error> {
        Ok(serde_json::from_str(contents.as_ref())?)
    }

    fn deserialize(&self, cfg: Config) -> Result<String, Error> {
        Ok(serde_json::to_string(&cfg)?)
    }

    pub fn load(&mut self) -> Config {
        let default = Config {
            twitter: self.default_twitter.clone(),
            users: HashMap::new(),
        };
        match self.read() {
            Ok(contents) => contents,
            _ => default,
        }
    }

    pub fn save(&mut self, cfg: Config) -> Result<(), Error> {
        Ok(self.write(cfg)?)
    }

    pub fn write(&mut self, cfg: Config) -> Result<(), Error> {
        let contents = self.deserialize(cfg)?;
        self.file.set_len(0)?;
        self.file.seek(SeekFrom::Start(0))?;
        Ok(self.file.write_all(contents.as_bytes())?)
    }
}
