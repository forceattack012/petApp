import { Component, OnInit } from '@angular/core';
import { Dashboard } from 'src/app/models/dashboard';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';
import { DomSanitizer } from '@angular/platform-browser';
import { Router } from '@angular/router';

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

  constructor(private petService: PetService, private sanitizer: DomSanitizer, private router: Router) {
   }

  ngOnInit(): void {
    this.loadData();
  }

  ngOnChanges(){
    this.id = '';
    this.isEdit = false;
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
    // this.router.navigate(['edit', id, 'edit']);
  }

  remove(id : string) {
    this.petService.remove(id).subscribe(b => {
      window.location.reload();
    })
  }
}
