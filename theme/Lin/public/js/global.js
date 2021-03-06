/*
 * Global js
 */
$(document)
.ready(function() {
  // create sidebar and attach to menu open
  $('.ui.sidebar')
    .sidebar('attach events', '.toc.item')
  ;
  
  $('.list-container')
    .transition('slide right in')
  ;

  $('.side-container')
    .transition('slide down in')
  ;
  
  $('.ui.sticky')
    .sticky({
      context: '.content-container',
      pushing: true
    })
  ;
});