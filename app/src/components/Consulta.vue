<template>

    <div>
        <b-form-select
        :options="recetas"
        v-model="escogido"
        :value="null"
        v-on:change="consultar"
        id="select"
        >
            <option slot="first" :value="null">Seleccione una receta</option>
        </b-form-select>

        <h2 v-if="estado=='consulta'">Consultar Receta</h2>
        <h2 v-if="estado=='actualizar'">Actualizar Receta</h2>

        <div v-if="estado=='consulta'" id="card" class="card" >
    
            <h3 class="card-text">{{nombre.toUpperCase()}}</h3>
            <h4 v-if="ingredientes.length!=0">Ingredientes: </h4>
            <ul class="list-group list-group-flush">
                <li class="list-group-item" v-for="(ingrediente,index) in ingredientes" v-bind:key="index">
                    {{index+1}}) {{ingrediente.cantidad}} {{ingrediente.unidad}} de {{ingrediente.nombre}}  
                    
                </li>
            </ul>
            <div class="card-body">
                <h4>Preparacion: </h4>
               <p>{{preparacion}}</p>
            </div>
            <div class="card-footer text-muted" id="foot">
                <b-button variant="outline-primary" v-on:click="estado='actualizar'" id="actualizar" >Actualizar</b-button>
                <b-button variant="outline-primary" v-on:click="borrarReceta" id="borrar" >Borrar</b-button>
            </div>
        </div>

        
        <div  v-if="estado=='actualizar'">
            <Crear :tipo="'actualizar'" :nom="nombre" :prep="preparacion" 
            :ing="ingredientes" v-on:cancelar="()=>{estado='null',escogido='null'}" 
            v-on:actualizado="(x)=>{consultarNombres();escogido=x;consultar(x)}"> </Crear>
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
                alert("Hubo un problema al conectar con la base de datos")
            });
        },

        consultar:function(event){
            console.log(event)

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
                if(response.body.ingredientes!=null) this.ingredientes = response.body.ingredientes
                else  this.ingredientes=[]
                
            }, response => {
                alert("Hubo un problema al conectar con la base de datos")
            });
            console.log(this.nombre)
            console.log(this.preparacion)
        },

        borrarReceta:function(){
            this.$http.delete('http://localhost:3000/receta/'+this.nombre)
            .then(response => {
                this.estado="null"
                this.escogido='null',
                this.consultarNombres()
                alert("Receta borrada exitosamente")
                this.estado='null'
            }, response => {
                alert("Hubo un problema al conectar con la base de datos")
            });            

        },
    }
}
   
</script>

<style>

#select{
    margin-top: 2%;
    margin-left: 25%;
    margin-right: 25%;
    width: 50%;
}

#card{
    
    width: 50%;
    margin-left: 25%;
    margin-right: 25%;
    margin-top: 2%;
    padding: 3%;
    border-radius: 10px;
    background-color: #ffec61;

}

#foot{
    background-color: #ffec61;
    text-align: center    
}

h2 {
    text-align: center;
    margin-top: 2%;
}

#actualizar{
    margin-right: 5%
}

</style>

