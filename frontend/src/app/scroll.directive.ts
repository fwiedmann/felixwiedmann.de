import {Directive, ElementRef, Input} from '@angular/core';
import {fromEvent, throttleTime} from "rxjs";

export type Config = {
  minMarginTop: number,
  maxMarginTop: number
  initMarginTop: number,
  updateStepsInPx: number
}

@Directive({
  selector: '[appScroll]'
})

export class ScrollDirective {
  currentMargin = 0;
  currentScroll = 0;

  constructor(
    private elementRef: ElementRef
  ) {
  }

  @Input() set appScroll(config: Config) {
    // only for large screens
    if (window.outerWidth < 1008) {
      return
    }

    this.currentMargin = config.initMarginTop;
    this.elementRef.nativeElement.style.marginTop = this.currentMargin + 'px'


    fromEvent(window, 'scroll').pipe(
      throttleTime(10)
    ).subscribe(event => {
      this.caluclateNewMargin(window.scrollY, config)
      this.elementRef.nativeElement.style.marginTop = this.currentMargin + 'px'
    });
  }

  private caluclateNewMargin(latestScroll: number, config: Config) {
    if (this.currentScroll === 0) {
      this.currentScroll = latestScroll
    }
    // scrolls down
    if (this.currentScroll < latestScroll) {
      this.currentScroll = latestScroll

      if ((this.currentMargin - config.updateStepsInPx) < config.minMarginTop) {
        return
      }

      this.currentMargin -= config.updateStepsInPx

    }
    // scrolls up
    if (this.currentScroll > latestScroll) {
      this.currentScroll = latestScroll

      if ((this.currentMargin + config.updateStepsInPx) > config.maxMarginTop) {
        return
      }
      this.currentMargin += config.updateStepsInPx

    }
  }
}
