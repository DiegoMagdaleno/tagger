//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "libtags/FileProperties.h"

// TFFileProperties represents an struct of the file properties
typedef struct _TGFileProperties {
    NSString *name;
    NSArray *tags;
} TGFileProperties;

const TGFileProperties*
filePropertiesData(FileProperties *file){
    TGFileProperties *filePropertiesData = malloc(sizeof(TGFileProperties));
    filePropertiesData->name = file.name;
    filePropertiesData->tags = file.tags;
    return filePropertiesData;
}

// Returns raw C string value 
const char*
NSStringToCString(NSString *s) {
    if (s == NULL) {return NULL; }

    const char *CStr = [s UTF8String];
    return CStr;
}

// Returns raw int value 
int
NSNumberToInt(NSNumber *i) {
    if (i == NULL ) {return 0;}
    return i.intValue;
}

// Returns lenght of array as long value 
unsigned long
NSArrayLen(NSArray *arrWithoutDetermined) {
    if (arrWithoutDetermined == NULL) { return 0; }

    return arrWithoutDetermined.count;
}

// Helper to move from C array to Go arrays
const void*
NSArrayItem(NSArray* arrWithDetermined, unsigned long i) {
    if (arrWithDetermined== NULL) { return NULL; }

    return [arrWithDetermined objectAtIndex:i];
}


