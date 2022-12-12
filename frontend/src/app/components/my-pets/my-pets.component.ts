import { OwnerService } from 'src/app/services/owner.service';
import { Component, OnInit } from '@angular/core';
import { Mypet } from 'src/app/models/mypet';
import { UserService } from 'src/app/services/user.service';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-my-pets',
  templateUrl: './my-pets.component.html',
  styleUrls: ['./my-pets.component.css']
})
export class MyPetsComponent implements OnInit {
  myPets: Mypet[] = [];

  constructor(private ownerService: OwnerService, private userService: UserService, private sanitizer: DomSanitizer) { }

  ngOnInit(): void {
    this.ownerService.getOwners(this.userService.getUserId()).subscribe(result => {
      this.myPets = result
      console.log(this.myPets)
    })
  }

  ngOnChanges(): void {
  }

  b64Image(base64: string) {
    return this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${base64}`);
  }

  remove(id: string) {
    this.ownerService.deleOwner(id).subscribe(result => {
      window.location.reload();
    })
  }

}
