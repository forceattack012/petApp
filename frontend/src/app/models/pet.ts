import { FilePet } from "./file";

export class Pet {
    ID : string = '';
    name: string = '';
    type: string = '';
    description: string = '';
    age: number = 0;
    files: Array<FilePet> = [];
}
