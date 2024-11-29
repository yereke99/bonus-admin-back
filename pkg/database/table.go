package database

var (
	// Таблица business_centers
	businessCentersTable = `
	CREATE TABLE IF NOT EXISTS business_centers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		published BOOLEAN DEFAULT NULL,
		city VARCHAR(255) DEFAULT NULL,
		name VARCHAR(255) DEFAULT NULL,
		address VARCHAR(255) DEFAULT NULL,
		hours VARCHAR(255) DEFAULT NULL,
		type VARCHAR(255) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица company_saved_barcodes
	companySavedBarcodesTable = `
	CREATE TABLE IF NOT EXISTS company_saved_barcodes (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
		code VARCHAR(255) UNIQUE NOT NULL,
		price BIGINT DEFAULT NULL,
		percent INT DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (company_id, code)
	);`

	// Таблица transactions
	transactionsTable = `
	CREATE TABLE IF NOT EXISTS transactions (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		company_id UUID REFERENCES companies(id) ON DELETE SET NULL,
		company_asset_id UUID REFERENCES company_assets(id) ON DELETE SET NULL,
		customer_id UUID REFERENCES customers(id) ON DELETE SET NULL,
		super_admin_id UUID REFERENCES super_admins(id) ON DELETE SET NULL,
		operator_type VARCHAR(255) NOT NULL,
		operator_id UUID NOT NULL,
		type VARCHAR(255) NOT NULL CHECK (type IN ('add', 'spend')),
		amount INT NOT NULL,
		comment TEXT DEFAULT NULL,
		system_comment TEXT DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица customers
	customersTable = `
	CREATE TABLE IF NOT EXISTS customers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		qr VARCHAR(255) UNIQUE NOT NULL,
		disabled BOOLEAN DEFAULT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) DEFAULT NULL,
		language VARCHAR(255) DEFAULT NULL,
		country VARCHAR(255) DEFAULT NULL,
		city VARCHAR(255) DEFAULT NULL,
		bonus BIGINT DEFAULT NULL,
		first_name VARCHAR(255) DEFAULT NULL,
		last_name VARCHAR(255) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица companies
	companiesTable = `
	CREATE TABLE IF NOT EXISTS companies (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		disabled BOOLEAN DEFAULT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		login VARCHAR(255) UNIQUE DEFAULT NULL,
		password VARCHAR(255) DEFAULT NULL,
		country VARCHAR(255) DEFAULT NULL,
		city VARCHAR(255) DEFAULT NULL,
		language VARCHAR(255) DEFAULT NULL,
		bonus BIGINT DEFAULT NULL,
		name VARCHAR(255) DEFAULT NULL,
		brand_name VARCHAR(255) DEFAULT NULL,
		iin VARCHAR(255) DEFAULT NULL,
		address VARCHAR(255) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица company_assets
	companyAssetsTable = `
	CREATE TABLE IF NOT EXISTS company_assets (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		disabled BOOLEAN DEFAULT NULL,
		published BOOLEAN DEFAULT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		login VARCHAR(255) UNIQUE DEFAULT NULL,
		password VARCHAR(255) DEFAULT NULL,
		company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
		business_center_id UUID REFERENCES business_centers(id) ON DELETE SET NULL,
		language VARCHAR(255) DEFAULT NULL,
		country VARCHAR(255) DEFAULT NULL,
		city VARCHAR(255) DEFAULT NULL,
		name VARCHAR(255) DEFAULT NULL,
		brand_name VARCHAR(255) DEFAULT NULL,
		iin VARCHAR(255) DEFAULT NULL,
		address VARCHAR(255) DEFAULT NULL,
		hours VARCHAR(255) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица super_admins
	superAdminsTable = `
	CREATE TABLE IF NOT EXISTS super_admins (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		disabled BOOLEAN DEFAULT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Таблица email_otps
	emailOtpsTable = `
	CREATE TABLE IF NOT EXISTS email_otps (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		email VARCHAR(255) UNIQUE NOT NULL,
		otp VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
)
