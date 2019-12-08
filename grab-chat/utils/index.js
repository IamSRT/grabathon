export const GetNestedValue = (object, propertyPath = '', defaultValue) => {
    // if there is no prop path OR object then return object if present otherwise defaultValue
    if (!propertyPath.length || !object) return object || defaultValue;
    // split propPath and reduce using nesting function
    const resultValue = propertyPath
        .split('.')
        .reduce((a, b, i, arr) => (arr.length - 1 === i ? a[b] : a[b] || {}), object);
    // return result value if available otherwise default Value
    return resultValue || defaultValue;
};

export const isEmpty = value => {
    switch (typeof value) {
        case 'string':
            return value.length === 0;
        case 'number':
        case 'boolean':
            return !value;
        case 'undefined':
            return true;
        case 'object':
            return !!(value === null || Object.keys(value).length === 0); // handling for null.
        case 'function':
            return false;
        default:
            return !value;
    }
};

export const ParseJSON = (data) => {
    try {
        return JSON.parse(data)
    } catch(e) {
        return null
    }
}