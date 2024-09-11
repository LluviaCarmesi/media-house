import type IDropdownOption from "../interfaces/IDropdownOption";

export default function createDropdownOptions(options: string[]) {
    let dropdownOptions: IDropdownOption[] = [];
    for (let i = 0; i < options.length; i++) {  
        dropdownOptions.push({
            value: options[i],
            label: options[i]
        });
    }
    
    return dropdownOptions;
}