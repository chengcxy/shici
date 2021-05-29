import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const store = () =>
  new Vuex.Store({
    state: {
      counter: 0,
      navMenuList: [],
      dynastyMenu: "",
      poeterMenu: "",
      poemMenu: ""
    },
    mutations: {
      increment(state) {
        state.counter++;
      },
      updateMenuList(state, value) {
        state.navMenuList = value;
        state.dynastyMenu = {
          title: value[0].dynasty,
          linkPath: "/dynasty/" + value[0].dynasty_id
        };

        state.poeterMenu = "";
        state.poemMenu = "";
      },
      changeCrumbsList(state, value) {
        console.log("stateValue");
        console.log(value);
        if (value.pageType == "dynasty") {
          state.dynastyMenu = {
            title: value.dynasty,
            linkPath: "/dynasty/" + value.dynasty_id
          };
          state.poeterMenu = "";
          state.poemMenu = "";
        } else if (value.pageType == "poeter") {
          state.poeterMenu = {
            title: value.poeter_name,
            linkPath: "/poeter/" + value.poeter_id
          };
          state.poemMenu = "";
        } else {
          state.poemMenu = {
            title: value.poem_name,
            linkPath: "/poem/" + value.poem_name
          };
        }
      }
    }
  });

export default store;
