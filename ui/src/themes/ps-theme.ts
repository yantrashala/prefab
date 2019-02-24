import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import { TextField } from '@material-ui/core';

const theme = createMuiTheme({
  palette: {
    primary: { main: '#ff9999' },
    secondary: { main: '#11cb5f' }
  },
  overrides: {
    MuiFormControl: {
      root: {
        marginNormal: { main: '45px' }
      }
    }
  }
});
export default theme;
