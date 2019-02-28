import React, { useState } from 'react';
import Button from '@material-ui/core/Button';
import { MuiThemeProvider, withStyles } from '@material-ui/core/styles';
import psTheme from '../../themes/ps-theme';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import '../../App.css';

import FormLabel from '@material-ui/core/FormLabel';
import FormControl from '@material-ui/core/FormControl';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormHelperText from '@material-ui/core/FormHelperText';
import Checkbox from '@material-ui/core/Checkbox';

const hostingProviders = [
  {
    value: 'aws',
    label: 'AWS'
  },
  {
    value: 'gcp',
    label: 'Google Cloud Platform'
  },
  { value: 'azure', label: 'Azure' }
];

const codeReviews = [
  {
    value: 'crucible',
    label: 'Crucible'
  },
  {
    value: 'sonar',
    label: 'Sonar Qube'
  }
];
const CiCdSetup = (props: any) => {
  const { classes } = props;
  const [hostingProvider, setHostingProvider] = useState('aws');
  const [logging, setLogging] = useState('elk');
  const [qualityReport, setQualityReport] = useState('speedy');
  const [codeReview, setCodeReview] = useState('crucible');


  const onLoggingChange = (event:any) => setLogging(event.target.value);
  const onQualityReportChange = (event:any) => setQualityReport(event.target.value);
  

  return (
    <MuiThemeProvider theme={psTheme}>
      <h1>CI-CD Setup</h1>
      <FormControl className="form-container">
        <TextField
          id="hosting-provider"
          select
          fullWidth
          label="Select your Hosting Provider"
          value={hostingProvider}
          onChange={() => setHostingProvider(hostingProvider)}
          margin="normal"
        >
          {hostingProviders.map(option => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>

        <FormLabel>Select the Environements you need</FormLabel>
        <FormGroup row>
          <FormControlLabel control={<Checkbox value="Develop" />} label="Develop" />
          <FormControlLabel control={<Checkbox value="Staging" />} label="Staging" />
          <FormControlLabel control={<Checkbox value="Production" />} label="Production" />
        </FormGroup>
        
        <FormGroup row />
       
        <FormLabel>Select the logging Framework</FormLabel>
        <RadioGroup
            aria-label="Logging"
            name="logging"
            value={logging}
            onChange={onLoggingChange}
          >
           <FormControlLabel value="elk" control={<Radio />} label="ELK" />
           <FormControlLabel value="efk" control={<Radio />} label="EFK" />

          </RadioGroup>

        <FormGroup row />
        <FormLabel>Location for your Engineering Quality Results</FormLabel>
        
        <RadioGroup aria-label="quality-reports" name="quality-reports" value={qualityReport} onChange={onQualityReportChange}>
        
            <FormControlLabel value="speedy" control={<Radio />} label="Speedy" />
            <FormControlLabel value="s3" control={<Radio />} label="S3" />
        
        </RadioGroup>


        <TextField
          id="code-reviews"
          select
          fullWidth
          label="Select Code Review Tools"
          value={codeReview}
          onChange={() => setCodeReview(codeReview)}
          margin="normal"
        >
          {codeReviews.map(option => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <Button variant="contained" color="primary">
          Setup Pipeline{' '}
        </Button>
      </FormControl>
    </MuiThemeProvider>
  );
};
export default CiCdSetup;
