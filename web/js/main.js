$(document).ready(function(){
  /**************************** */
  //side menu arrow transformation
  /**************************** */ 
  $(".side-main-nav-link").click(function(){
     if ($(this).find(".open-arrow").hasClass("active")) {
          $(this).find(".open-arrow").removeClass("active");
      } else {
          $(this).find(".open-arrow").addClass("active");
      } 
  });
  /**************************** */
  //side menu toggle function
  /**************************** */ 
  $(".side-menu-toggle").click(function(){
    $(".side-bar").toggleClass("active");
    $(".main-content").toggleClass("full");
  });
  
  //convert title to slug
  $('.input-slug').slugify('.input-title');
});