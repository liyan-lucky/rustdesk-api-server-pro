export type Certificate = {
    key: string;
    cert: string;
};
export type CertificateAuthorityOptions = {
    organization: string;
    countryCode: string;
    state: string;
    locality: string;
    validity: number;
};
export declare function createCA(options: CertificateAuthorityOptions): Promise<Certificate>;
export type CertificateOptions = {
    domains: string[];
    validity: number;
    organization?: string;
    email?: string;
    ca: Certificate;
};
export declare function createCert(options: CertificateOptions): Promise<Certificate>;
