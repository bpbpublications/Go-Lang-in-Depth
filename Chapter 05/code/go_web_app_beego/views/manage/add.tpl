<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Add A Product</h1>
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    
    <h2>Product Details</h2>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}


    
    <p>          
      <form role="form" id="user" method="POST">
        <div class="form-group {{if .Errors.Name}}has-error has-feedback{{end}}">
          <label for="name">Product name： {{if .Errors.Name}}({{.Errors.Name}}){{end}}</label>
          <input name="name" type="text" value="{{.Product.Name}}" class="form-control" tabindex="1" />
          {{if .Errors.Name}}<span class="glyphicon glyphicon-remove form-control-feedback"></span>{{end}}
        </div>
        
        <div class="form-group">
          <label for="category">Category：</label>
          <input name="category" type="text" value="{{.Product.Category}}" class="form-control" tabindex="2" />
        </div>
        
        <div class="form-group">
          <label for="image">Image：</label>
          <input name="image" type="text" value="{{.Product.Image}}" class="form-control" tabindex="3" />
        </div>
        
        <input type="submit" value="Create Product" class="btn btn-default" tabindex="4" /> &nbsp;
        <a href="#" title="don't create the product">Cancel</a>
      </form>
    </p>
  </div>
</div>