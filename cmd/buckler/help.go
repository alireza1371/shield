package main

import (
	fmt "github.com/jhunt/go-ansi"
)

func ShowHelp(command string) {
	switch (command) {
	case "api": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} api [OPTIONS] @Y{ALIAS} @Y{URL}\n")
		fmt.Printf("\n")
		fmt.Printf("  Target a SHIELD Core.\n")
		fmt.Printf("\n")
		fmt.Printf("  Connects to the given SHIELD URL (i.e. @B{https://shield.example.com}),\n")
		fmt.Printf("  validates that it is healthy and trustworthy, and then saves the\n")
		fmt.Printf("  connection details in your @M{$SHIELD_CONFIG} or @M{--config}.\n")
		fmt.Printf("\n")
		fmt.Printf("  From then on, that SHIELD Core can be contacted by setting the\n")
		fmt.Printf("  @M{$SHIELD_CORE} environment variable to the supplied @M{ALIAS},\n")
		fmt.Printf("  or by passing @M{--config ALIAS} on the command-line, explicitly.\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Options:}\n")
		fmt.Printf("\n")
		fmt.Printf("  -k, --skip-ssl-validate   Do not validate the SSL/TLS certificate\n")
		fmt.Printf("                            presented by the remote SHIELD Core.\n")
		fmt.Printf("                            @R{THIS IS HIGHLY DISCOURAGED.}\n")
		fmt.Printf("\n")
		fmt.Printf("  --ca-cert ...             Provide either a path to a CA certificate\n")
		fmt.Printf("  --ca-certificate ...      file, or the literal PEM data of an X.509\n")
		fmt.Printf("                            certificate representing the Certificate\n")
		fmt.Printf("                            Authority that signed the SHIELD Core's TLS\n")
		fmt.Printf("                            certificate.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "cores": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} cores\n")
		fmt.Printf("\n")
		fmt.Printf("  Lists all known SHIELD Cores.\n")
		fmt.Printf("\n")
		fmt.Printf("  Consults your @M{$SHIELD_CONFIG} (or, if given, @M{--config}), and\n")
		fmt.Printf("  prints out pertinent information about each targeted core.\n")
		fmt.Printf("  This includes the URL, the alias, and any security parameters.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "curl": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} curl [@Y{METHOD}] @Y{RELATIVE-URL} [@Y{BODY}]\n")
		fmt.Printf("\n")
		fmt.Printf("  Issues raw HTTP requests to the targeted SHIELD Core.\n")
		fmt.Printf("\n")
		fmt.Printf("  Any valid HTTP method can be provided, but buckler will default to\n")
		fmt.Printf("  using @M{GET}.  Some HTTP methods, like POST and PUT, may require\n")
		fmt.Printf("  that you supply a request body, like the JSON for updating a target\n")
		fmt.Printf("  data system.\n")
		fmt.Printf("\n")
		fmt.Printf("  If you are using buckler curl to debug and/or troubleshoot your\n")
		fmt.Printf("  SHIELD Core, keep in mind that you may need to add @M{--trace} to your\n")
		fmt.Printf("  invocation if you want to see the HTTP requests and headers.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "id": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} id\n")
		fmt.Printf("\n")
		fmt.Printf("  Displays information about your currently authenticated session.\n")
		fmt.Printf("\n")
		fmt.Printf("  This includes your profile information (display name, username,\n")
		fmt.Printf("  etc.), as well as what system role you have been assigned, and\n")
		fmt.Printf("  what tenants (if any) you have been granted membership into.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "init": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} init [--master @Y{PASSWORD}]\n")
		fmt.Printf("\n")
		fmt.Printf("  Initializes a brand new SHIELD Core.\n")
		fmt.Printf("\n")
		fmt.Printf("  SHIELD maintains an internal encrypted vault of secrets,\n")
		fmt.Printf("  for protecting your data archives with strong encryption.\n")
		fmt.Printf("  Before SHIELD will be able to use this vault, it needs to\n")
		fmt.Printf("  be initialized, and a SHIELD master password chosen.\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Options:}\n")
		fmt.Printf("\n")
		fmt.Printf("  --master ...    New master password. (@W{$SHIELD_CORE_MASTER})\n")
		fmt.Printf("\n")
		fmt.Printf("  In general, use of the command-line flag is discouraged.\n")
		fmt.Printf("  Instead, for automation purposes, set @W{$SHIELD_CORE_MASTER}\n")
		fmt.Printf("  in your environment.\n")
		fmt.Printf("\n")
		fmt.Printf("  For interactive use, not specifying @M{--master} causes buckler\n")
		fmt.Printf("  to prompt you for a new master password, with appropriate\n")
		fmt.Printf("  security precautions (no terminal echo, confirmation, etc.)\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "login": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} login [--username @Y{USERNAME}] [--password @Y{PASSWORD}]\n")
		fmt.Printf("       @G{buckler} login --token @Y{AUTH-TOKEN}\n")
		fmt.Printf("       @G{buckler} login --providers\n")
		fmt.Printf("       @G{buckler} login --via @Y{PROVIDER-ID}\n")
		fmt.Printf("\n")
		fmt.Printf("  Authenticate to a SHIELD Core.\n")
		fmt.Printf("\n")
		fmt.Printf("  There are three ways to authenticate to a SHIELD Core: local user\n")
		fmt.Printf("  authentication, auth tokens, and 3rd-party provider authentication.\n")
		fmt.Printf("  Each way has its own set of options, and its own applications.\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Local User Authentication:}\n")
		fmt.Printf("\n")
		fmt.Printf("  -u, --username ...    Who to log in as. (@W{$SHIELD_CORE_USERNAME})\n")
		fmt.Printf("  -p, --password ...    Secret password.  (@W{$SHIELD_CORE_PASSWORD})\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Token Authentication:}\n")
		fmt.Printf("\n")
		fmt.Printf("      --token ...       Auth Token to use. (@W{$SHIELD_CORE_TOKEN})\n")
		fmt.Printf("\n")
		fmt.Printf("@B{3rd-Party Authentication}\n")
		fmt.Printf("\n")
		fmt.Printf("      --providers       List the 3rd-party authentication providers\n")
		fmt.Printf("                        that have been configured by the SHIELD site\n")
		fmt.Printf("                        operators, and are usable by @B{--via}.\n")
		fmt.Printf("\n")
		fmt.Printf("      --via ...         Which 3rd-party authentication provider to\n")
		fmt.Printf("                        use for logging in.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "logout": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} logout\n")
		fmt.Printf("\n")
		fmt.Printf("  Logs you out of your active, authenticated session.\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "rekey": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} rekey [--old-master @Y{PASSWORD}]\n")
		fmt.Printf("                     [--new-master @Y{PASSWORD}]\n")
		fmt.Printf("\n")
		fmt.Printf("  Change your SHIELD Master Password.\n")
		fmt.Printf("\n")
		fmt.Printf("  SHIELD maintains an internal encrypted vault of secrets,\n")
		fmt.Printf("  for protecting your data archives with strong encryption.\n")
		fmt.Printf("  The keys to this vault are protected by a master password.\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Options:}\n")
		fmt.Printf("\n")
		fmt.Printf("  --old-master ...    Current master password.\n")
		fmt.Printf("  --new-master ...    New master password.\n")
		fmt.Printf("\n")
		fmt.Printf("  In general, use of the command-line flags is discouraged.\n")
		fmt.Printf("\n")
		fmt.Printf("  Rekeying should, in general, be an interactive process.\n")
		fmt.Printf("  Not specifying @M{--master} causes buckler to prompt you\n")
		fmt.Printf("  for both the current master password, and your desired new\n")
		fmt.Printf("  master password, with appropriate security precautions (no\n")
		fmt.Printf("  terminal echo, confirmation, etc.)\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */
	case "unlock": /* {{{ */
		fmt.Printf("USAGE: @G{buckler} unlock [--master @Y{PASSWORD}]\n")
		fmt.Printf("\n")
		fmt.Printf("  Unlock a SHIELD Core.\n")
		fmt.Printf("\n")
		fmt.Printf("  SHIELD maintains an internal encrypted vault of secrets,\n")
		fmt.Printf("  for protecting your data archives with strong encryption.\n")
		fmt.Printf("  Before SHIELD will be able to use this vault, it needs to\n")
		fmt.Printf("  be unlocked, using the SHIELD master password.\n")
		fmt.Printf("\n")
		fmt.Printf("@B{Options:}\n")
		fmt.Printf("\n")
		fmt.Printf("  --master ...    Master password. (@W{$SHIELD_CORE_MASTER})\n")
		fmt.Printf("\n")
		fmt.Printf("  In general, use of the command-line flag is discouraged.\n")
		fmt.Printf("  Instead, for automation purposes, set @W{$SHIELD_CORE_MASTER}\n")
		fmt.Printf("  in your environment.\n")
		fmt.Printf("\n")
		fmt.Printf("  For interactive use, not specifying @M{--master} causes buckler\n")
		fmt.Printf("  to prompt you for the master password, with appropriate\n")
		fmt.Printf("  security precautions (no terminal echo)\n")
		fmt.Printf("\n")
		fmt.Printf("\n")

	/* }}} */

	default:
		fmt.Printf("No help is available for `@G{buckler} @R{%s}` yet.\n\n", command)
	}
}
