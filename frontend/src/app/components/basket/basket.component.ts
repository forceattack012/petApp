import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Basket } from 'src/app/models/basket';
import { BasketService } from 'src/app/services/basket.service';
import { UserService } from 'src/app/services/user.service';
import { DomSanitizer } from '@angular/platform-browser';
import { Pet } from 'src/app/models/pet';
import { Owner } from 'src/app/models/onwer';
import { OwnerService } from 'src/app/services/owner.service';

@Component({
  selector: 'app-basket',
  templateUrl: './basket.component.html',
  styleUrls: ['./basket.component.css']
})
export class BasketComponent implements OnInit {
  @Input()
  isBasket = false;
  @Input()
  basket: Basket;
  username = '';

  @Output()
  pets = new EventEmitter();
  @Output()
  removePetId = new EventEmitter();

  owners: Owner[] = []

  constructor(private basketService: BasketService, private userService: UserService,
    private sanitizer: DomSanitizer, private ownerService: OwnerService) { }

  ngOnInit(): void {
  }

  ngOnChanges(): void {
  }

  b64Image(base64: string) {
    return this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${base64}`);
  }

  remove(id : string){
    this.basket.pets = this.basket.pets.filter(r => r.ID !== id)
    this.pets.emit(this.basket.pets)
    this.removePetId.emit(id);

    this.basketService.addBasket(this.basket).subscribe(p => {
      this.basketService.getBasket(this.basket.name).subscribe(basket => {
        this.basket = basket;
      })
    })
  }

  submit(items: any){
    items.forEach(item =>{
      let owner = {
        user_id:  parseInt(this.userService.getUserId()),
        pet_id: item.ID
      }
      this.owners.push(owner)
    })
    console.log(this.owners)
    this.ownerService.addOwners(this.owners).subscribe(p =>{
      window.location.href = "/"
    })
  }

}
