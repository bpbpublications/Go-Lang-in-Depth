<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Products List</h1>
     {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Name</th>
          <th>Category</th>
          <th>Image</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Id}}</td>
          <td>{{$record.Name}}</td>
          <td>{{$record.Category}}</td>
          <td>{{$record.Image}} {{urlfor "ManageController.Delete" ":id" "21"}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="4"><a href="{{urlfor "ManageController.Add"}}" title="add new product">Add another product</a></td>
        </tr>
      </tfoot>
    </table>
  </div>
</div>