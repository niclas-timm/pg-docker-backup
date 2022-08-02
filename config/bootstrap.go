package config

// BootstrapConfig loads the configuration from the yaml files
// and stores them in the Conf object.
func BootstrapConfig(){
	error := ParseYmlFiles(&Conf, "config.default.yml", "config.yml")

	if error != nil {
		panic("Could not load config file.")
	}
}