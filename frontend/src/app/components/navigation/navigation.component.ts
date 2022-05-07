import {Component, OnInit} from '@angular/core';
import {animate, state, style, transition, trigger} from "@angular/animations";

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss'],
  animations: [
    trigger('openClose', [
      state("closed", style({
        width: '0%',
      })),
      state("open", style({
        width: '100%',
      })),
      transition('open => closed', [
        animate('0.2s')
      ]),
      transition('closed => open', [
        animate('0.2s')
      ]),
    ])
  ]
})
export class NavigationComponent implements OnInit {
  open: boolean = false

  constructor() {
  }

  ngOnInit(): void {
  }

  scroll(target: string) {

    const el = document.getElementById(target);
    this.open = false
    if (!el) {
      return
    }
    el.scrollIntoView({behavior: 'smooth'});
  }
}
