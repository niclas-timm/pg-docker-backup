package config

const (
	TmpDirName = "tmp"
	TmpImpDirName = "tmp/imp"
	TmpImpFileName = "tmp_import_file_sql.gz"
)

// Global config object.
// Populated with data from config.yml files after BootstrapConfig is executed
var  Conf = Config{}
