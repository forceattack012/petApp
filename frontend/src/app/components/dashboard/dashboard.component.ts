import { Component, Input, OnInit, Output } from '@angular/core';
import { Dashboard } from 'src/app/models/dashboard';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';
import { DomSanitizer } from '@angular/platform-browser';
import { Basket } from 'src/app/models/basket';
import { UserService } from 'src/app/services/user.service';
import { BasketService } from 'src/app/services/basket.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  items : Dashboard;
  page : number = 1;
  limit: number = 12;
  totalRows = 1;
  isEdit = false;
  id ='';
  isBasket = false;

  @Input()
  pets = [];
  basket: Basket = new Basket();

  @Output()
  removePetId = ''

  constructor(private petService: PetService, private sanitizer: DomSanitizer,
    private userService: UserService, private basketService: BasketService
    ) {
   }

  ngOnInit(): void {
    this.loadData();
  }

  ngOnChanges(){
    this.id = '';
    this.isEdit = false;
  }

  ngOnDestroy(){
    this.isEdit = !this.isBasket;
    this.basket = new Basket();
  }

  loadPage(page: number) {
    console.log(page)
    this.page = page;
    this.loadData();
  }

  loadData() {
    this.petService.getPets(this.page, this.limit).subscribe(p =>{
      this.items = p;
      this.totalRows = p.total_rows;
    })
  }

  b64Image(base64: string) {
    return this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${base64}`);
  }

  gotoEdit(id: string){
    this.id = id;
    this.isEdit = true;
  }

  add(items: Pet){
    if(this.pets.length >= 4){
      alert("Add pets only 4 pets")
      return;
    }

    document.getElementById("add-"+items.ID).setAttribute('disabled', 'true')
    document.getElementById("edit-"+items.ID).setAttribute('disabled', 'true')
    document.getElementById("remove-"+items.ID).setAttribute('disabled', 'true')

    this.pets.push(items);
    this.basket.name = this.userService.getName();
    this.basket.pets = this.pets;

    this.basketService.addBasket(this.basket).subscribe(p =>{
      this.isBasket = true;
      let username = this.userService.getName();
      this.basketService.getBasket(username).subscribe(p =>{
        this.basket = new Basket();
        this.basket = p;
      })
    });
  }

  remove(id : string) {
    this.petService.remove(id).subscribe(b => {
      window.location.reload();
    })
  }

  onChangePet(event : any) {
    console.log(event)
    this.pets = event;
  }

  onChangeRemoveId(id: string){
    console.log('removeId =>' +id)
    document.getElementById("add-"+id).removeAttribute('disabled')
    document.getElementById("edit-"+id).removeAttribute('disabled')
    document.getElementById("remove-"+id).removeAttribute('disabled')
  }
}
