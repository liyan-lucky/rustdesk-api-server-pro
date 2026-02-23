#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = require("commander");
const node_fs_1 = require("node:fs");
const promises_1 = require("node:fs/promises");
const node_path_1 = require("node:path");
const mkcert_1 = require("./mkcert");
commander_1.program
    .command("create-ca")
    .option("--organization [value]", "organization name", "Test CA")
    .option("--country-code [value]", "country code", "US")
    .option("--state [value]", "state name", "California")
    .option("--locality [value]", "locality address", "San Francisco")
    .addOption(new commander_1.Option("--validity [days]", "validity in days")
    .default(365)
    .argParser((val) => Number.parseInt(val, 10)))
    .option("--key [file]", "output key file", "ca.key")
    .option("--cert [file]", "output certificate file", "ca.crt")
    .action(async (options) => {
    const ca = await (0, mkcert_1.createCA)(options);
    await (0, promises_1.writeFile)(options.key, ca.key);
    console.log(`CA Private Key: ${options.key}`);
    await (0, promises_1.writeFile)(options.cert, ca.cert);
    console.log(`CA Certificate: ${options.cert}`);
});
commander_1.program
    .command("create-cert")
    .alias("create-certificate")
    .option("--ca-key [file]", "ca private key file", "ca.key")
    .option("--ca-cert [file]", "ca certificate file", "ca.crt")
    .addOption(new commander_1.Option("--validity [days]", "validity in days")
    .default(365)
    .argParser((val) => Number.parseInt(val, 10)))
    .option("--key [file]", "output key file", "cert.key")
    .option("--cert [file]", "output certificate file", "cert.crt")
    .option("--organization [value]", "optional organization name")
    .option("--email [value]", "optional email address")
    .option("--domains, --domain [values...]", "domains or ip addresses", ["localhost", "127.0.0.1"])
    .action(async (options) => {
    let ca = {
        key: await (0, promises_1.readFile)(options.caKey, "utf-8").catch(() => void 0),
        cert: await (0, promises_1.readFile)(options.caCert, "utf-8").catch(() => void 0)
    };
    if (!ca.key || !ca.cert) {
        console.error("Unable to find CA key or certificate.");
        console.error("Please run `mkcert create-ca` to create a new certificate authority.");
        return;
    }
    const cert = await (0, mkcert_1.createCert)({
        ca: { key: ca.key, cert: ca.cert },
        domains: options.domain,
        validity: options.validity,
        organization: options.organization,
        email: options.email
    });
    await (0, promises_1.writeFile)(options.key, cert.key);
    console.log(`Private Key: ${options.key}`);
    await (0, promises_1.writeFile)(options.cert, `${cert.cert}${ca.cert}`); // write full chain certificate
    console.log(`Certificate: ${options.cert}`);
});
function getVersion() {
    return JSON.parse((0, node_fs_1.readFileSync)((0, node_path_1.resolve)(__dirname, "../package.json"), "utf-8")).version;
}
commander_1.program.version(getVersion()).parseAsync(process.argv);
