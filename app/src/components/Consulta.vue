<template>

    <div>
        <b-form-select
        :options="recetas"
        v-model="escogido"
        :value="null"
        v-on:change="consultar"
        >
            <option slot="first" :value="null">Seleccione una receta</option>
        </b-form-select>

        <div v-if="estado=='consulta'" class="card" style="width: 18rem;">
    
            <h3 class="card-text">{{nombre.toUpperCase()}}</h3>
            <h4>Ingredientes: </h4>
            <ul class="list-group list-group-flush">
                <li class="list-group-item" v-for="(ingrediente,index) in ingredientes" v-bind:key="index">
                    {{index+1}}) {{ingrediente.cantidad}}.{{ingrediente.unidad}} de {{ingrediente.nombre}}  
                    
                </li>
            </ul>
            <div class="card-body">
                <h4>Preparacion: </h4>
               <p>{{preparacion}}</p>
            </div>
            <div cclass="card-footer text-muted">
                <b-button variant="outline-primary" v-on:click="estado='actualizar'" id="actualizar" >Actualizar</b-button>
                <b-button variant="outline-primary" v-on:click="borrarReceta" id="borrar" >Borrar</b-button>
            </div>
        </div>

        
        <div  v-if="estado=='actualizar'">
            <Crear :tipo="'actualizar'" :nom="nombre" :prep="preparacion" 
            :ing="ingredientes" v-on:cancelar="()=>{estado='null',escogido='null'}"> </Crear>
        </div>

    </div>
    
</template>

<script>

import Ingrediente from './Ingrediente'
import Crear from  './Crear'

export default {
    name: 'consulta', 
    components: {
        Ingrediente,
        Crear
    },
    data () {
        return{
            estado:'null',
            escogido:'null',
            recetas:[],
            nombre:'',
            preparacion:'',
            ingredientes:[]
        }
    },

    created: function () {
        this.consultarNombres()
    },

    methods: {
        consultarNombres: function(){
            this.$http.get('http://localhost:3000/receta')
            .then(response => {
                this.recetas=response.body
            }, response => {
                alert("No, sirvio!!! :(")
            });
        },

        consultar:function(event){

            if(event==null){
                this.estado="null"
                return
            }            

            this.$http.get('http://localhost:3000/receta/'+event)
            .then(response => {
                this.nombreInicial=response.body.nombre
                this.estado='consulta'
                this.nombre=response.body.nombre
                this.preparacion=response.body.preparacion
                this.ingredientes = response.body.ingredientes 
                
            }, response => {
                
            });
        },

        borrarReceta:function(){
            this.$http.delete('http://localhost:3000/receta/'+this.nombre)
            .then(response => {
                this.estado="null"
                this.escogido='null',
                this.consultarNombres()
                alert("si, sirvio!!! :(")
                this.estado='null'
            }, response => {
                alert("No, sirvio!!! :(")
            });            

        },
    }
}
   
</script>

<style>

</style>

