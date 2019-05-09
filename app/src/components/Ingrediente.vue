<template>
    <div id="ingre">
        <b-row>
            <b-col>
                <b-input
                type="text"
                v-model="nombre"
                placeholder="Nombre Ingrediente"
                v-on:keyup="$emit('childToParent', {nom:nombre, ident:num})"
                required
                ></b-input>
            </b-col>

            <b-col >
                <b-input
                type="number"
                step="0.1"
                v-model="cantidad"
                placeholder="cantidad"
                v-on:keyup="$emit('childToParent', {cant:parseFloat(cantidad), ident:num})"
                v-on:change="$emit('childToParent', {cant:parseFloat(cantidad), ident:num})"
                required
                ></b-input>
            </b-col>
            
            <b-col >
                <b-form-select
                :options="{ 'Ltr': 'Ltr', 'gr': 'gr', 'kilo': 'kilo' }"
                v-model="unidad"
                :value="null"
                v-on:change="$emit('childToParent', {unid:unidad, ident:num})"
                >
                    <option slot="first" :value="null">Elegir unidad de medida</option>
                </b-form-select>
            </b-col>
            <b-col >
                <b-button variant="outline-primary" id="crear" v-on:click="$emit('eliminar',num)">Eliminar</b-button>
            </b-col>
        </b-row>
    </div>
</template>

<script>

export default {
    name: 'Crear', 
    data () {
        return{
            estado:'inicio',
            unidad:'null',
            nombre:'',
            cantidad:''
        }
    },

    props:{
        num:Number,
        nom: {
            type: String,
            default: ''
        },
        uni: {
            type: String,
            default: 'null'
        },
        can: {
            type: Number,
            default: 0
        }
    },

    created: function(){
        
        this.unidad=this.uni,
        this.nombre=this.nom,
        this.cantidad=this.can
        
    }    
}
</script>
<style>

#ingre{
    margin-bottom: 5%    
}

</style>