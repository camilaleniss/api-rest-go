<template>
  <div class="mt-5">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <h1>{{msg}}</h1>
  <div class="d-flex justify-content-center">
  <form>
    <label for="exampleInputDomain">Enter a domain: </label>
    <input type="text" v-model="domainSearched" class="form-control" placeholder="domain.com" aria-describedby="emailHelp">
    <br>
    <button type="submit" v-on:click="getDomain" class="btn btn-primary">Search domain</button>
    <br>
    <br>
    <button type="submit" v-on:click="getDomains" class="btn btn-primary">Historial</button>
  </form>
  </div>
  <br>
  <div class="d-flex justify-content-center">

  <div class="row">
  <div class="col-sm-6">
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">Historial</h5>
        <ul>
        <p v-for="domain in domains" :key="domain.host">
          {{ domain.host }}
        </p>
        </ul>
      </div>
    </div>
  </div>
  <div class="col-sm-6">
    <div class="card">
      <br>
      <div class="card-body">
        <h5 class="card-title">{{domainSearched}}</h5>
        <p class="card-subtitle mb-2 text-muted">Servers Info: </p>
        <p v-for="server in servers" :key="server.address">
          IP address: {{ server.address }} <br>
          SSL Grade: {{server.ssl_grade}}<br>
          Country: {{server.country}}<br>
          Owner: {{server.owner}}<br>
        </p>
        <br>
        <p class="card-subtitle mb-2 text-muted">Domain Info: </p>
        <p class="card-text">Servers changed: {{ssl_change}}</p>
        <p class="card-text">SSL grade: {{ssl_grade}}</p>
        <p class="card-text">SSL previous grade: {{ssl_previous_grade}}</p>
        <p class="card-text">Logo: {{logo}}</p>
        <p class="card-text">Title: {{title}}</p>
        <p class="card-text">Is down: {{is_down}}</p>
      </div>
    </div>
  </div>
</div>
  </div>
  </div>

  

</template>

<script>

import axios from 'axios';

export default {
  name: 'DomainApp',

  props: {
    msg: String,
  },

  data(){
    return {
      domainSearched: "",
      domains: null,
      servers : null,
      ssl_grade: null,
      ssl_previous_grade : null,
      ssl_change: null,
      title: null,
      logo: null,
      is_down: null
    }
  },



  methods: {
    getDomains: function(e){
      e.preventDefault();
      axios.get('http://localhost:8082/api/domains').then(response =>
      {
        console.log(response.data);
        this.domains = response.data.items;
      }, error => {
        console.log(error);
      }
      );


    },
    getDomain: function(e){
      e.preventDefault();
      axios.get('http://localhost:8082/api/'+this.domainSearched).then(response =>
      {
        console.log(response.data);
        this.servers = response.data.servers;
        this.ssl_change = response.data.servers_changed;
        this.ssl_previous_grade = response.data.previous_ssl_grade;
        this.logo = response.data.logo;
        this.title = response.data.title;
        this.is_down = response.data.is_down;
        
      }, error => {
        console.log(error);
      }
      );
    }

  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
