"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createCert = exports.createCA = void 0;
const node_forge_1 = require("node-forge");
const node_net_1 = __importDefault(require("node:net"));
const node_util_1 = require("node:util");
async function generateCert(options) {
    const { subject, issuer, extensions, validity } = options;
    const generateKeyPair = (0, node_util_1.promisify)(node_forge_1.pki.rsa.generateKeyPair.bind(node_forge_1.pki.rsa));
    // create random serial number between between 50000 and 99999
    const serial = Math.floor(Math.random() * 95000 + 50000).toString();
    const keyPair = await generateKeyPair({ bits: 2048, workers: 4 });
    const cert = node_forge_1.pki.createCertificate();
    // serial number must be hex encoded
    cert.serialNumber = Buffer.from(serial).toString("hex");
    cert.publicKey = keyPair.publicKey;
    cert.setSubject(subject);
    cert.setIssuer(issuer);
    cert.setExtensions(extensions);
    cert.validity.notBefore = new Date();
    cert.validity.notAfter = new Date();
    cert.validity.notAfter.setDate(cert.validity.notAfter.getDate() + validity);
    // sign the certificate with it's own
    // private key if no separate signing key is provided
    const signWith = options.signWith ? node_forge_1.pki.privateKeyFromPem(options.signWith) : keyPair.privateKey;
    cert.sign(signWith, node_forge_1.md.sha256.create());
    return {
        key: node_forge_1.pki.privateKeyToPem(keyPair.privateKey),
        cert: node_forge_1.pki.certificateToPem(cert)
    };
}
async function createCA(options) {
    // certificate Attributes: https://git.io/fptna
    const attributes = [
        { name: "commonName", value: options.organization },
        { name: "countryName", value: options.countryCode },
        { name: "stateOrProvinceName", value: options.state },
        { name: "localityName", value: options.locality },
        { name: "organizationName", value: options.organization }
    ];
    // required certificate extensions for a certificate authority
    const extensions = [
        { name: "basicConstraints", cA: true, critical: true },
        { name: "keyUsage", keyCertSign: true, critical: true }
    ];
    return await generateCert({
        subject: attributes,
        issuer: attributes,
        extensions: extensions,
        validity: options.validity
    });
}
exports.createCA = createCA;
async function createCert(options) {
    // certificate Attributes: https://git.io/fptna
    const attributes = [
        { name: "commonName", value: options.domains[0] } // use the first address as common name
    ];
    if (options.organization) {
        attributes.push({ name: "organizationName", value: options.organization });
    }
    if (options.email) {
        attributes.push({ name: "emailAddress", value: options.email });
    }
    // required certificate extensions for a tls certificate
    const extensions = [
        { name: "basicConstraints", cA: false, critical: true },
        {
            name: "keyUsage",
            digitalSignature: true,
            keyEncipherment: true,
            critical: true
        },
        { name: "extKeyUsage", serverAuth: true, clientAuth: true },
        {
            name: "subjectAltName",
            altNames: options.domains.map((domain) => {
                // types https://git.io/fptng
                const TYPE_DOMAIN = 2;
                const TYPE_IP = 7;
                if (node_net_1.default.isIP(domain)) {
                    return { type: TYPE_IP, ip: domain };
                }
                return { type: TYPE_DOMAIN, value: domain };
            })
        }
    ];
    const ca = node_forge_1.pki.certificateFromPem(options.ca.cert);
    return await generateCert({
        subject: attributes,
        issuer: ca.subject.attributes,
        extensions: extensions,
        validity: options.validity,
        signWith: options.ca.key
    });
}
exports.createCert = createCert;
