/* Do not change, this code is generated from Golang structs */


export class User {
    id: string;
    email: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.email = source["email"];
    }
}