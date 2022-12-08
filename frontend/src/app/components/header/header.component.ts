import { Component, OnInit } from '@angular/core';
import { NgbNavConfig } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  constructor(config: NgbNavConfig) {
    config.destroyOnHide = false;
		config.roles = false;
  }

  ngOnInit(): void {
  }

}
