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
        
        <div v-if="escogido!='null' && escogido!=null" class="card" style="width: 18rem;">
    
            <h3 class="card-text">{{nombre.toUpperCase()}}</h3>
            <h4>Ingredientes: </h4>
            <ul class="list-group list-group-flush">
                <li class="list-group-item" v-for="(ingrediente,index) in ingredientes" v-bind:key="index">
                    {{index+1}}. {{ingrediente.nombre}}: {{ingrediente.cantidad}} .{{ingrediente.unidad}}
                    
                </li>
            </ul>
            <div class="card-body">
                <h4>Preparacion: </h4>
               <p>{{preparacion}}</p>
            </div>
            <div cclass="card-footer text-muted">
                <b-button variant="outline-primary" id="actualizar" >Actualizar</b-button>
                <b-button variant="outline-primary" id="consultar" >Consultar</b-button>
            </div>
        </div>

    </div>
    
</template>

<script>

export default {
    name: 'Crear', 
    data () {
        return{
            estado:'null',
            escogido:'null',
            recetas:[],
            nombre:'',
            preparacion:'',
            ingredientes:{}
        }
    },

    created: function () {
        this.$http.get('http://localhost:3000/receta')
        .then(response => {
            this.recetas=response.body
        }, response => {
            alert("No, sirvio!!! :(")
        });

    },

    methods: {
        consultar:function(event){

            if(event==null){
                this.estado='null'
                return
            }
            

            this.$http.get('http://localhost:3000/receta/'+event)
            .then(response => {
                console.log(response.body.ingredientes)
                this.nombre=response.body.nombre
                this.preparacion=response.body.preparacion
                this.ingredientes = response.body.ingredientes/*.map((x)=>{
                   return [x.nombre,x.cantidad,x.unidad]
                })*/
                console.log(this.ingredientes)
                
            }, response => {
                
            });
        }
    }
}
   
</script>

<style>

</style>

