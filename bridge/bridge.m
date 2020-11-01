//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "ext/NSTaggerURL.h"
#include "ext/FileProperties.h"

#include <Foundation/Foundation.h>


// FileProoperties is an NSObject, describes a file
// name -> path to the file
// tags -> Name of the tags that are contained
@interface FileProperties : NSObject
@property (nonatomic, strong) NSString *name;
@property (nonatomic, strong) NSArray *tags;

- (id)initWithName:(NSString *)name tags:(NSArray *)tags;

@end

@implementation FileProperties
@synthesize name;
@synthesize tags;

- (id)initWithName:(NSString *)name tags:(NSArray *)tags
{
    if (self = [super init]) {
        self.name = name;
        self.tags = tags;
    }
    return self;
}
@end

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

// Lists all files with tags at a collection of paths, that are later converted
// into NSURL
NSArray* listFilesWithTagsAtNSURL(NSArray* path) {
    NSMutableArray *filePreliminary = [NSMutableArray array];

    for (id eachPathAsURL in path) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];

        if([target GetTags] != nil) {
            [filePreliminary addObject:[target path]];
        }
    } 
    return filePreliminary;
}

// Converts a given dictionary to a FileProperties object
NSArray*
fromDictionaryToFilePropertiesArray(NSDictionary *targetDict) {
    NSMutableArray *objectList = [NSMutableArray array];
    
    
    for (id key in targetDict) {
        FileProperties *target = [[FileProperties alloc] init];
        target.name = key;
        target.tags = [targetDict objectForKey:key];
        [objectList addObject:target];
    }
    
    return objectList;
}

const NSArray*
getFilesWithCertainMacOSTag(char const*  pathRaw) {
    
    NSString *path = [NSString stringWithUTF8String:pathRaw];

    NSMutableDictionary *mappedFiles = [NSMutableDictionary dictionary];

    BOOL isDir;
    NSArray *folders;

    NSFileManager *fm = [NSFileManager defaultManager];

    [fm fileExistsAtPath:path isDirectory:&isDir];


    /*
     * Stage 1:
     * We are now on ObjC so we need to convert our Path string to
     * an NSURL
     */
    NSURL *directoryURL = [[NSURL alloc] initFileURLWithPath:path];

    if (!isDir) {
        folders = @[directoryURL];
    } else {
        folders = [fm contentsOfDirectoryAtURL:directoryURL includingPropertiesForKeys:NULL options:NSDirectoryEnumerationSkipsHiddenFiles error:NULL];
    }

    
    for (id eachPathAsURL in folders) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];
        
        mappedFiles[[target path]] = [target GetTags];        
    }
    
    return fromDictionaryToFilePropertiesArray(mappedFiles);
    
}
