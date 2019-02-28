import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import { TextField } from '@material-ui/core';

const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: { main: '#fe414d' },
    secondary: { main: '#11cb5f' }
  },

  //-- Still need to get the component specific overides to work
  overrides: {
    MuiInput: {
      root: {
        width: '450px',
        marginNormal: { main: '45px' }
      }
    },

    MuiFormControl: {
      root: {
        marginTop: '15px',
        marginBottom: '15px'
      }
    }
  }
});
export default theme;
