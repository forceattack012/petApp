import { Pet } from "./pet";

export class Dashboard{
    limit: number;
    page: number;
    total_rows: number;
    total_pages: number;
    items: Array<Pet> = [];
}
